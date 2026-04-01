-- Criar banco manualmente no pdAdmin -> CREAT BATABASE IF NOT EXISTS devbook;
-- USE devbook;

-- Remove a tabela se ela já existir para evitar erros em re-execuções
DROP TABLE IF EXISTS usuarios;

-- Criação da tabela com sintaxe nativa PostgreSQL
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL, -- Recomendado 100 para armazenar hashes
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Comentário opcional para documentação na interface do pgAdmin
COMMENT ON TABLE usuarios IS 'Tabela para armazenar dados dos usuários do sistema devbook';