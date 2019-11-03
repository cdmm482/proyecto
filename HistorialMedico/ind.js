var firebaseConfig = {
  apiKey: "AIzaSyADzH0A_V_rBPtLzt0-ibqJPJOxjqmw514",
  authDomain: "sis22019.firebaseapp.com",
  databaseURL: "https://sis22019.firebaseio.com",
  projectId: "sis22019",
  storageBucket: "sis22019.appspot.com",
  messagingSenderId: "652240941407",
  appId: "1:652240941407:web:8a924537d8be9af0a7c6ea",
  measurementId: "G-RZXJ3ZLDQZ"
};
firebase.initializeApp(firebaseConfig);

firebase.auth().onAuthStateChanged(function (user) {
  if (user) {
    // User is signed in.

    document.getElementById("user_div").style.display = "block";
    document.getElementById("login_div").style.display = "none";

    var user = firebase.auth().currentUser;

    if (user != null) {

      var email_id = user.email;
      document.getElementById("user_para").innerHTML = "Welcome User : " + email_id;
    }

  } else {
    // No user is signed in.
    // window.location = "menu.html";
    document.getElementById("user_div").style.display = "none";
    document.getElementById("login_div").style.display = "block";

  }
});

function login() {

  var userEmail = document.getElementById("email_field").value;
  var userPass = document.getElementById("password_field").value;

  firebase.auth().signInWithEmailAndPassword(userEmail, userPass).catch(function (error) {
    // Handle Errors here.
    var errorCode = error.code;
    var errorMessage = error.message;

    window.alert("Error : " + errorMessage);

    // ...
  });

}

function logout() {
  firebase.auth().signOut();
}
