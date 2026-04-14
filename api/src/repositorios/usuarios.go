package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values ($1, $2, $3, $4)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return erro
	}

	return nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeouNick string) ([]modelos.Usuario, error) {
	nomeouNick = fmt.Sprintf("%%%s%%", nomeouNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criado_em from usuarios where nome like $1 or nick like $2",
		nomeouNick, nomeouNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuario do banco de dados
func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criado_em from usuarios where id = $1",
		ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuario no banco de dados
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = $1, nick = $2, email = $3 where id = $4",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informacoes de um usuario no banco de dados
func (repositorio usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = $1",
	)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuario pelo email e retorna seu id e senha com hash
func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = $1", email,
	)

	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir permite que um usuario siga outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert into seguidores (usuario_id, seguidor_id) values ($1, $2) on conflict (usuario_id, seguidor_id) do nothing",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// PararDeSeguir permite que um usuario pare de seguir o outro
func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = $1 and seguidor_id = $2",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores traz todos os seguidores de um usuario
func (reposotorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := reposotorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criado_em 
		  from usuarios u 
		 inner join seguidores s 
		    on s.seguidor_id = u.id 
		 where s.usuario_id = $1`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var seguidores []modelos.Usuario

	for linhas.Next() {
		var seguidor modelos.Usuario

		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		seguidores = append(seguidores, seguidor)

	}

	return seguidores, nil

}

// BuscarSeguidores traz todos os usuarios que um determinado usario está seguindo
func (reposotorio usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {

	linhas, erro := reposotorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criado_em 
		  from usuarios u 
		 inner join seguidores s 
		    on s.usuario_id = u.id 
		 where s.seguidor_id = $1`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)

	}

	return usuarios, nil

}

// BuscarSenha traz a senha de um usuario pelo ID
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = $1", usuarioID)
	if erro != nil {
		return "", erro
	}

	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuario
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = $1 where id = $2")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
