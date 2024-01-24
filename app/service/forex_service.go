package service

import (
	"fmt"

	"github.com/xiaoma/trading/app/model"
	"github.com/xiaoma/trading/platform/database"
)

func CreateTrading(trading *model.ForexTradingModel) error {
	db, _ := database.OpenDBConnection()
	err := db.Create(&trading).Error
	return err
}

func GetTradings() ([]model.ForexTradingModel, error) {
	var tradingList []model.ForexTradingModel
	db, _ := database.OpenDBConnection()
	sql := "SELECT * FROM forex_trading_models"
	res := db.Raw(sql).Scan(&tradingList)

	if res.RowsAffected == 0 {
		return make([]model.ForexTradingModel, 0), nil
	}
	return tradingList, nil
}

// update role
func UpdateForex(trading *model.ForexTradingModel) error {
	db, _ := database.OpenDBConnection()
	err := db.Save(&trading).Error
	return err
}

func DeleteForex(tradingID uint64) error {
	db, _ := database.OpenDBConnection()
	err := db.Where("id = ?", tradingID).Delete(&model.ForexTradingModel{}).Error
	return err
}

func GetTradingByID(tradingID uint64) (model.ForexTradingModel, error) {
	fmt.Println(tradingID)
	db, _ := database.OpenDBConnection()
	trading := model.ForexTradingModel{}
	err := db.Where(&model.ForexTradingModel{ID: tradingID}).First(&trading).Error
	if err != nil {
		return trading, err
	}
	return trading, nil

}
