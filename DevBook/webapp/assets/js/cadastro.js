$("#formulario-cadastro").on('submit', criarUsuario);

function criarUsuario(evento) {
  evento.preventDefault()

  if ($('#senha').val() != $('#confirmar-senha').val()) {
    Swal.fire("Erro! ", "As senhas não coincidem", "erro")
    return
  }

  $.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val(),
    }
  }).done(function () {

    Swal.fire("Sucesso!", "Usuário cadastrado com sucesso", "success").then(function () {
      $.ajax(
        {
          url: "/login",
          method: "post",
          data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
          }
        }
      ).done(function () {
        
        location.href= "/home"
      }).fail(function () {
        Swal.fire("Erro! ", "Erro ao logar", "erro")
      })
    })

  }).fail(function (erro) {
    console.log(erro)
    $('#nome').trigger("focus")
    Swal.fire("Erro! ", "Erro ao cadastrar o usuário", "erro")
  })
}