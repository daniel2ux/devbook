$('#login').on('submit', doLogin);

function doLogin(e) {
    e.preventDefault();

    $.ajax({
        url: '/login',
        method: 'POST',
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        },
    })
        .done(function () {
            window.location = '/home';
        })
        .fail(function () {
            alert('Login failed');
        });
}
