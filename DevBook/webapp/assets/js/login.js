$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {
  evento.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      senha: $('#senha').val()
    },
    complete: callback
  });

  function callback(response) {
    if (response.status == 200) {
      window.location = "/home"
    } else {
      Swal.fire("Erro! ", "Usuario ou senha inválidos", "error")
    }
  }
}

