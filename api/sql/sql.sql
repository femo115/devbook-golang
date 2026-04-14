-- Criar banco manualmente no pdAdmin -> CREAT BATABASE IF NOT EXISTS devbook;
-- USE devbook;

-- Remove a tabela se ela já existir para evitar erros em re-execuções
DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
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

CREATE TABLE seguidores (
    usuario_id INT NOT NULL,
    seguidor_id INT NOT NULL,
    
    -- Chaves Estrangeiras (Foreign Keys)
    CONSTRAINT fk_usuario 
        FOREIGN KEY (usuario_id) 
        REFERENCES usuarios(id) 
        ON DELETE CASCADE,
        
    CONSTRAINT fk_seguidor 
        FOREIGN KEY (seguidor_id) 
        REFERENCES usuarios(id) 
        ON DELETE CASCADE,

    -- Chave Primária Composta
    PRIMARY KEY (usuario_id, seguidor_id)
);

-- Comentário opcional para documentação na interface do pgAdmin
COMMENT ON TABLE seguidores IS 'Tabela para armazenar dados dos seguidores de um usuario do sistema devbook';

CREATE TABLE publicacoes (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(50) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    
    autor_id INT NOT NULL,
    CONSTRAINT fk_autor
        FOREIGN KEY (autor_id)
        REFERENCES usuarios(id)
        ON DELETE CASCADE,

    curtidas INT DEFAULT 0,
    criada_em TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE publicacoes IS 'Tabela para armazenar as publicacoes dos usuarios';
