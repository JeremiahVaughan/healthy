package models

type UnexpectedErrorsModel struct {

}

func NewUnexpectedErrorsModel() *UnexpectedErrorsModel {
    return &UnexpectedErrorsModel{
    }
}

func (u *UnexpectedErrorsModel) Internal(err error) {
    // todo implement
}
