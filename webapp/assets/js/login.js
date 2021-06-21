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
			toastr.options = {
				"progressBar": true,
				"timeOut": "3000",
			};
			toastr.error('Access Unauthorized');
		});
}
