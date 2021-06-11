$('#new-post').on('submit', newPost);

function newPost(e) {
    e.preventDefault();

    $.ajax({
        url: '/posts',
        method: 'POST',
        data: {
            title: $('#title').val(),
            post: $('#post').val(),
        },
    })
        .done(function () {
            window.location = '/home';
        })
        .fail(function () {
            alert('Publish fail');
        });
}
