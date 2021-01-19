INSERT INTO usuarios (nome, nick, email, senha)
VALUES 
("Belchior Pereira de Araújo Júnior", "Belchior Júnior", "belchiorjr@gmail.com", "$2a$10$ajhVn6Q/Aw5bKOfiSvhunuKiMVE.Agc1MqKXP.jkRNjh7EUjUeREq"),
("Taiane Lopes de Lima Araújo", "Taiane", "taiane@gmail.com", "$2a$10$ajhVn6Q/Aw5bKOfiSvhunuKiMVE.Agc1MqKXP.jkRNjh7EUjUeREq"),
("Ana Laura Lima Araújo", "Ana Laura", "analaura@gmail.com", "$2a$10$ajhVn6Q/Aw5bKOfiSvhunuKiMVE.Agc1MqKXP.jkRNjh7EUjUeREq"),
("Filipe Lima Araújo", "Filipe", "filipe@gmail.com", "$2a$10$ajhVn6Q/Aw5bKOfiSvhunuKiMVE.Agc1MqKXP.jkRNjh7EUjUeREq");

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES (1,2),(1,3),(1,4),(2,1),(2,3),(2,4),(3,1),(3,2),(3,4),(4,1),(4,2),(4,3);


INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES  
("Publicação do Usuário 1", "Essa é a publicação do usuário 1! Oba!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2! Oba!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3! Oba!", 3),
("Publicação do Usuário 3", "Essa é a publicação do usuário 4! Oba!", 4);