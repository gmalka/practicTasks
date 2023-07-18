all			:
			docker-compose up myapp

test		:	cli_test server_test

cli_test	:
			cd ./cli && go test

server_test	:
			docker-compose up test

clean		:
			docker-compose down