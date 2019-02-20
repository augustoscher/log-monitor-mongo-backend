package model

//LogGroup model representing grouped logs
type LogGroup struct {
	CodigoIntegracao string `bson:"codigo_integracao" json:"codigo_integracao"`
	NomeIntegracao   string `bson:"nome_integracao" json:"nome_integracao"`
	TipoNotificacao  string `bson:"tipo_notificacao" json:"tipo_notificacao"`
	Quantidade       int64  `bson:"quantidade" json:"quantidade"`
}
