$('#form-createUser').on('submit', createUser);

function createUser(event){

    event.preventDefault();

    password = $("#password").val();
    confirmPassword = $("#confirmPassword").val();

    if (password != confirmPassword){
        Swal.fire('Atenção!', 'Senha é diferente do confirmar senha.', 'warning')
        return
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $("#name").val(),
            nickname: $("#nickname").val(),
            email: $("#email").val(),
            password: $("#password").val(),
        }
    //status code 2xx
    }).done(function(){
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
        .then(function() {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    password: $('#password').val()
                }
            }).done(function() {
                window.location = "/home";
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
            })
        })
    //status code 4xx/5xx
    }).fail(function(errorMsg){
        Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
    })


}