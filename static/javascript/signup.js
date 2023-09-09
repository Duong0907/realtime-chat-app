var signupBtn = document.getElementById('signup-btn');

function signup() {
    return fetch('/api/user/signup', {
        method: 'POST',
        body: JSON.stringify({
            'username': document.getElementById('username-input').value,
            'password': document.getElementById('password-input').value,
            'email': document.getElementById('email-input').value
        })
    }).then(res => res.json())
}

signupBtn.onclick = function() {
    signup()
        .then(() => {
            window.location.href = '/login';
        })
}