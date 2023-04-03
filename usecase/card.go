package usecase

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm/utils"
	"shopChallenge/domain"
	"strconv"
	"strings"
	"time"
)

func (u UseCaseImpl) Transfer(
	req domain.TransferRequest) error {
	txn := domain.Transactions{
		CardID:   req.SourceCardNo,
		ToCardID: req.TargetCardNo,
		Amount:   req.Amount,
		IsActive: true,
	}

	txnRules, err := u.Repo.ReadTransactionRules("transfer")
	if err != nil {
		return err
	}

	fa := &finalAmount{}
	fa.applyRules(txn.Amount, txnRules)
	txn.Amount = decimal.Decimal(*fa)

	err = u.ValidateTransfer(&txn, txnRules)
	if err != nil {
		return err
	}
	fmt.Println("b- s---->", txn.Card.Account)
	fmt.Println("b- t---->", txn.ToCard.Account)

	err = u.transfer(&TransferAccountRequest{
		Amount:    txn.Amount,
		SourceAcc: &txn.Card.Account,
		TargetAcc: &txn.ToCard.Account,
	})
	if err != nil {
		return err
	}
	fmt.Println("a- s---->", txn.Card.Account)
	fmt.Println("a- t---->", txn.ToCard.Account)

	err = u.sendMessageToCustomers(&txn, &txnRules)
	if err != nil {
		return err
	}

	err = u.InsertTransactions(txn)
	if err != nil {
		return err
	}

	return err
}

func (u UseCaseImpl) InsertTransactions(txn domain.Transactions) error {
	// from-account TXN
	err := u.Repo.InsertTransaction(txn)
	if err != nil {
		return err
	}

	// to-account TXN
	// note: in to-account TXN from-account will be to-account with negative amount
	// therefore all the TXNs of one account can be retrieved with from-account index
	txn.Card, txn.ToCard = txn.ToCard, txn.Card
	txn.CardID, txn.ToCardID = txn.ToCardID, txn.CardID
	txn.Amount = txn.Amount.Neg()
	err = u.Repo.InsertTransaction(txn)
	if err != nil {
		return err
	}
	return nil
}
func (u UseCaseImpl) ValidateTransfer(
	txn *domain.Transactions, txnRules domain.TransactionRules) error {

	toCard := card{Cards: domain.Cards{
		CardID: txn.ToCardID},
		rules:  txnRules,
		amount: txn.Amount,
	}
	err := u.ToCardValidation(&toCard)
	if err != nil {
		return err
	}
	txn.ToCard = &toCard.Cards

	fromCard := card{Cards: domain.Cards{
		CardID: txn.CardID},
		rules:  txnRules,
		amount: txn.Amount,
	}
	err = u.fromCardValidation(&fromCard)
	if err != nil {
		return err
	}
	txn.Card = &fromCard.Cards

	return nil
}

func (u UseCaseImpl) fromCardValidation(c *card) error {
	if !c.CheckCardNumber() {
		return domain.ErrNotValidFromCard
	}

	var err error
	ac := domain.AccountAccess{
		Contained: true,
		CustomerAccess: domain.CustomerAccess{
			Contained: true,
		},
	}
	c.Cards, err = u.Repo.ReadCard(c.CardID, ac)
	if err != nil {
		return err
	}

	fmt.Println("@@@ok error", c.amount, c.rules.MinLimit, c.rules.MaxLimit)
	if (!c.rules.MinLimit.IsZero() && c.amount.LessThan(c.rules.MinLimit)) ||
		(!c.rules.MaxLimit.IsZero() && c.amount.GreaterThan(c.rules.MaxLimit)) {
		fmt.Println("!!!!slimit error", c.amount, c.rules.MinLimit, c.rules.MaxLimit)
		return domain.ErrAmountNotInRange
	}
	return nil
}

func (u UseCaseImpl) ToCardValidation(card *card) error {
	fmt.Println("++++", card.CardID)
	if !card.CheckCardNumber() {
		return domain.ErrNotValidToCard
	}
	var err error
	ac := domain.AccountAccess{
		Contained: true,
		CustomerAccess: domain.CustomerAccess{
			Contained: true,
		},
	}
	card.Cards, err = u.Repo.ReadCard(card.CardID, ac)
	fmt.Println("&&&&", card.Cards)
	if err != nil {
		return err
	}
	return nil
}
func (u UseCaseImpl) transfer(t *TransferAccountRequest) error {
	t.SourceAcc.Balance = t.SourceAcc.Balance.Sub(t.Amount)

	if t.SourceAcc.Balance.LessThan(decimal.Zero) {
		return domain.ErrNotEnoughCredit
	}
	err := u.Repo.UpdateAccount(*t.SourceAcc)
	if err != nil {
		return err
	}

	t.TargetAcc.Balance = t.TargetAcc.Balance.Add(t.Amount)
	err = u.Repo.UpdateAccount(*t.TargetAcc)
	if err != nil {
		return err
	}
	return nil
}

func (u UseCaseImpl) sendMessageToCustomers(
	txn *domain.Transactions, rules *domain.TransactionRules) error {
	toCard := card{
		Cards:  *txn.ToCard,
		rules:  *rules,
		amount: txn.Amount,
	}
	message := toCard.prepareSMS(rules.TemplateSms)
	_ = u.Sms.SendMessage(message, toCard.Account.Customer.MobileNumber)
	//if err != nil {
	//	return err
	//}
	fmt.Println("_____________________________")
	fromCard := card{
		Cards:  *txn.ToCard,
		rules:  *rules,
		amount: txn.Amount.Neg(),
	}
	message = fromCard.prepareSMS(rules.TemplateSms)
	_ = u.Sms.SendMessage(message, fromCard.Account.Customer.MobileNumber)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (c card) prepareSMS(template string) string {
	fmt.Println("bf ", template)
	functionality := "deposit"
	if c.amount.LessThan(decimal.Zero) {
		functionality = "withdrawal"
	}
	replacer := strings.NewReplacer(
		"NAME", c.Account.Customer.Name,
		"FUNC", functionality,
		"AMOUNT", c.amount.String(),
		"BALANCE", c.Account.Balance.String(),
		"FEE", c.rules.Fee.String(),
		"DATE", time.Now().Format(time.RFC822Z),
	)
	fmt.Println("af ", replacer.Replace(template))
	return replacer.Replace(template)
}

func (f *finalAmount) applyRules(
	amount decimal.Decimal, txnRules domain.TransactionRules) {
	*f = finalAmount(decimal.Sum(amount, txnRules.Fee))
}

func (c card) CheckCardNumber() bool {
	cr := utils.ToString(c.CardID)
	fmt.Println("cr:", cr)
	if len(cr) != 16 {
		fmt.Println("false")
		return false
	}
	var cardTotal int64 = 0
	for i, ch := range cr {
		c, err := strconv.ParseInt(string(ch), 10, 8)
		if err != nil {
			fmt.Println("false2")
			return false
		}
		if i%2 == 0 {
			if c*2 > 9 {
				cardTotal = cardTotal + (c * 2) - 9
			} else {
				cardTotal = cardTotal + (c * 2)
			}
		} else {
			cardTotal += c
		}
	}
	fmt.Println("false3", cardTotal%10 == 0)
	return cardTotal%10 == 0
}