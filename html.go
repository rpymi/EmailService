<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <style>
        body {
            font-family: 'Helvetica', sans-serif;
            color: #fff;
            margin: 0px;
            padding: 0px;
            background-color: #00ac85;
            width: 100%;
            height: 100vh;
            display: flex;
            flex-direction: column;
            justify-content: center;
        }

        .app__container {
            width: 520px;
            display: flex;
            flex-direction: column;
            margin: auto;
        }

        h1 {
            text-align: center;
            font-size: 34px;
        }

        .app__container input,
        .app__container textarea {
            padding: 7px 15px;
            margin-bottom: 20px;
            border: 0px;
            outline: none;
        }


        button {
            margin: auto;
            background: #fff;
            text-align: center;
            padding: 7px 15px;
            font-size: 20px;
            text-decoration: none;
            color: #000;
            font-weight: 400;
            border: 0px;
            width: 100%;
            outline: none;
        }
    </style>
</head>
    <body>
        <div class="app__container">
            <h1>Send Emails Using GoLang</h1>
           
            <input id="name" class="form-control" placeholder="Name"/>

            <input id="to" class="form-control" placeholder="Email" />

            <textarea id="message" class="form-control" placeholder="Message"></textarea>
           
           <button id="email-button">Send Email</button>
        </div>
    </body>
    <script>
        const sendEmailButton = document.querySelector('#email-button');

        sendEmailButton.onclick = (event) => {
            const nameInput = document.querySelector('#name');
            const toEmailInput = document.querySelector('#to');
            const messageInput = document.querySelector('#message');

            if(!nameInput.value) {
                alert("Name is empty");
            } else if(!toEmailInput.value) {
                alert("Email is empty");
            }else if(!messageInput.value) {
                alert("Message is empty");
            } else {
                fetch(`/sendEmail`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': undefined
                    },
                    body: JSON.stringify({
                        name: nameInput.value.trim(),
                        email: toEmailInput.value.trim(),
                        message: messageInput.value.trim()
                    })
                })
                .then((resp) => resp.json())
                .then((result) => {
                    if(result.Code === 200) {
                        alert(result.Message);
                    } else {
                        alert(result.Message)
                    }
                })
                .catch((error) => {
                    alert(error.Message)
                });
            }
        };
    </script>
</html>