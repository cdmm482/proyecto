
// require("firebase/firestore");


// var firebaseConfig = {
//     apiKey: "AIzaSyADzH0A_V_rBPtLzt0-ibqJPJOxjqmw514",
//     authDomain: "sis22019.firebaseapp.com",
//     databaseURL: "https://sis22019.firebaseio.com",
//     projectId: "sis22019",
//     storageBucket: "sis22019.appspot.com",
//     messagingSenderId: "652240941407",
//     appId: "1:652240941407:web:8a924537d8be9af0a7c6ea",
//     measurementId: "G-RZXJ3ZLDQZ",
// };

// // var firebase;
// firebase.initializeApp(firebaseConfig);

// function rev() {
//     firebase.auth().onAuthStateChanged(function (user) {
//         if (user) {
//             // User is signed in.
//             var user = firebase.auth().currentUser;

//             if (user != null) {

//                 var email_id = user.email;
//                 document.getElementById("user_para").innerHTML = "Welcome User : " + email_id;
//             }

//         } else {
//             // No user is signed in.
//             window.location = "index.html";
//             document.getElementById("user_div").style.display = "none";
//             document.getElementById("login_div").style.display = "block";

//         }
//     });
// }

function rev() {
    firebase.auth().onAuthStateChanged(function (user) {
        if (user) {
            // User is signed in.
            var user = firebase.auth().currentUser;

            if (user == null) {
                window.location = "index.html";
            }

        } else {
            // No user is signed in.
            window.location = "index.html";
            document.getElementById("user_div").style.display = "none";
            document.getElementById("login_div").style.display = "block";

        }
    });
}
function logout() {
    firebase.auth().signOut();
}