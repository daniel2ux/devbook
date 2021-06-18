$('#login').on('submit', login);

function login(e) {
	e.preventDefault();

	$.ajax({
		url: '/login',
		method: 'POST',
		data: {
			email: $('#email').val(),
			password: $('#password').val(),
		},
	})
		.done(() => {
			window.location = '/home';
		})
		.fail(() => {
			swal({
				text: 'Invalid Access!',
				icon: 'error',
				buttons: false,
				timer: 2000,
			});
		});
}
