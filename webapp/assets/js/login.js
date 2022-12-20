$('#form-login').on('submit', login);

function login (event){

    event.preventDefault();

    const email = $("#email").val();
    const password = $("#password").val();

    if (email == "" || password == ""){
        Swal.fire("Ops...", "Email e senha são campos obrigatórios", "error");
        return
    }

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: email,
            password: password
        }
    }).done(function(){
        window.location = "/home";
    }).fail(function(){
        Swal.fire("Ops...", "Usuário ou senha incorretos!", "error");
    })

}

