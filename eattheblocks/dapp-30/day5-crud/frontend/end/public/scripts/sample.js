
// var database = firebase.initializeApp(config).database();
var commentsRef = firebase.database().ref('posts/' + "2142");
commentsRef.on('child_added', function (data) {
    console.log("Se agrego ");
    // addCommentElement(postElement, data.key, data.val().text, data.val().author);
});

commentsRef.on('child_changed', function (data) {
    window.location = "index.html";
    // console.log("Se cambio");
    // setCommentValues(postElement, data.key, data.val().text, data.val().author);
});

commentsRef.on('child_removed', function (data) {
    console.log("Se elimino");
    // deleteComment(postElement, data.key);
});


async function call() {

    //     alert("Fulano0");
    //     await delay(5000);
    //     alert("Fulano");
    //     await delay(5000);
    //     alert("Fulano2");

    //     // setTimeout(function () {
    //     //     alert("Etapa 1 completada");
    //     //     setTimeout(function () {
    //     //         alert("Etapa 2 completada");
    //     //         setTimeout(function () {
    //     //             alert("Etapa 3 completada");
    //     //             setTimeout(function () {
    //     //                 alert("Etapa 4 completada");

    //     //             }, 10000);
    //     //         }, 10000);
    //     //     }, 10000);
    //     // }, 10000);
    // }

    // function call() {
    // const firebase = require("firebase/database");
    // require("firebase/firestore");

   

    // var userId = firebase.auth().currentUser.uid;
    // var refe = database.ref("posts");

    var db = firebase.database();

    var leadsRef = db.ref('posts');
    leadsRef.on('value', function (snapshot) {
        snapshot.forEach(function (childSnapshot) {
            var childData = childSnapshot.val();
            // alert(snapshot.value);
            alert(childData.author);
        });
    });

    // leadsRef.on('child_added', function (snapshot) {
    //     snapshot.forEach(function (childSnapshot) {
    //         var childData = childSnapshot.val();
    //         alert(childData);
    //     });        //Do something with the data
    // });

    //    
    // var alldata = "";
    // const dbRef = database.ref();
    // const usersRef = dbRef.child('posts');
    // usersRef.on('value', function (snapshot) {
    //     snapshot.forEach(function (childSnapshot) {
    //         var childData = childSnapshot.val();
    //         alldata = alldata + childData;
    //     });
    // });

    // Import Admin SDK
    // var admin = require("firebase-admin");

    // Get a database reference to our posts
    // var db = firebase.database();
    // var ref = db.ref("server/saving-data/fireblog/posts");

    // // Attach an asynchronous callback to read the data at our posts reference
    // ref.on("posts", function (snapshot) {
    //     console.log(snapshot.val());
    // }, function (errorObject) {
    //     console.log("The read failed: " + errorObject.code);
    // });

    // var recentPostsRef = firebase.database().ref('posts').limitToLast(100);
    // alert(recentPostsRef);
}