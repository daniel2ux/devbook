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
			iziToast.error({
				title: 'Error',
				message: 'Invalid Access!',
				position: 'topRight',
			});
		});
}
