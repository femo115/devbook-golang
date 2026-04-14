package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositorio de publicacoes
type publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio publicacoes) Criar(publicacao modelos.Publicacao) error {
	statment, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values ($1, $2, $3)",
	)
	if erro != nil {
		return erro
	}

	defer statment.Close()

	if _, erro = statment.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorID traz uma unica publicacao no banco de dados
func (repositorio publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
		select p.*, u.nick 
		  from publicacoes p 
		 inner join usuarios u
		    on u.id = p.autor_id
		 where p.id = $1
	`, publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicacoes dos usuarios seguidos e também do proprio usuario que fez a requisicao
func (repositorio publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select distinct p.*, u.nick
		  from publicacoes p
		 inner join usuarios u
		    on u.id = p.autor_id
		 inner join seguidores s
		    on p.autor_id = s.usuario_id
		 where u.id = $1
		    or s.seguidor_id = $2
		 order by p.id desc
		`, usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicacao no banco de dados
func (repositorio publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = $1, conteudo = $2 where id = $3")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if  _, erro = statement.Exec(publicacao.Titulo,publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}
	
	return nil
}