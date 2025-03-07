# Forma Padrão de Transmitir e armazenar objetos JSON
# Especificação: Forma compacta e segura de representar a Claims classificando as como Json, assinados ou criptografados digitalmente
# Para que aconteça isso é necessario o JWS(Json Web Signature) e JWE(Json Web Encryption). Disponibilizando varios tipos de assinaturas ou criptografias

# Diferença Assinado(JWS) e Criptografado(JWE): 
## JWS: Segredo como o algoritmo HMAC ou um par de Chaves publico-privada como o RSA ou ECDSA verificando a integridade das reinividicacoes contidas neles.
## JWE: Garante o segredo das informações entre as partes 
# JWT se destaca pela segurança das informações ali contida e possivel verificar se o conteudo do token e valido apenas verificando o proprio conteudo  

# Onde são utilizados: 
## Autenticação: Principalmente pela facil utilização em diferentes dominios
## Autorização: Uma vez q o usuário for logado cada requisição ao servidor será adicionado um JWT permitindo seu acesso a rota serviços e recursos liberados atraves daquele token. 
## Troca de Informações: Forma segura de transmitir entre partes.
# Exemplo de JWT: 

# Header: Algoritmo & Tipo de Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
{
    "alg": "HS256", (Tipo de Algoritmo d Criptografia na Assinatura)
    "typ": "JWT" (Tipo de Token)
}

# Payload: Dados:                    eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
{
    "iss": "https://token.com.br", Emissor do token
    "sub": "1234567890", Indentificador único do usuário para quem o token foi emitido
    "name": "Nome Qualquer",
    "admin": true Priveligio de adm
}

# Verificador de Assinatura:         TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
HMACSHA256(
    base64UrlEncode(header) + "." +
    base64UrlEncode(payload),
    secret)
