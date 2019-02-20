package model

import "gopkg.in/mgo.v2/bson"

//LogMessage representa uma mensagem de log
type LogMessage struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	CodigoFilial string        `bson:"codigo_filial" json:"codigo_filial"`
	NomeFilial   string        `bson:"nome_filial" json:"nome_filial"`
	//TipoNotificacao - F: Filial | M: Matriz
	CodigoIntegracao     string `bson:"codigo_integracao" json:"codigo_integracao"`
	NomeIntegracao       string `bson:"nome_integracao" json:"nome_integracao"`
	TipoNotificacao      string `bson:"tipo_notificacao" json:"tipo_notificacao"`
	DescricaoErro        string `bson:"descricao_erro" json:"descricao_erro"`
	ConteudoMensagemErro string `bson:"conteudo_mensagem_erro" json:"conteudo_mensagem_erro"`
	DataHora             string `bson:"data_hora" json:"data_hora"`
}
