import 'package:flutter/material.dart';
import 'package:test1/page/qr_scan.dart';
void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'consulta',
      theme: ThemeData(

        primarySwatch:  Colors.blue,
      ),
      home: Puente(),
    );
  }
}
class Puente extends StatefulWidget{
  @override
  MiProgra  createState() => MiProgra();
}
class MiProgra extends State<Puente>{

  var correo = TextEditingController();
  var pasw = TextEditingController();
  var _cont=true;
  var  _validate=false;
  var  _valpas=false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text("Iniciar Sesion"),),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.start,
          children: <Widget>[
            TextField(
              decoration: InputDecoration(labelText: 'CI', hintText: 'ingrese carned de indentidad',

              ),
              controller: correo,
              keyboardType: TextInputType.number,


            ),
            TextField(
              obscureText: _cont,
              decoration: InputDecoration(labelText: 'Password',
                errorText: _valpas ? 'contraseña incorecta' : null,
                hintText: 'ingrese contraseña',
                suffixIcon: IconButton(icon: Icon(Icons.adjust), onPressed: () {
                  setState(() {
                    if(_cont){
                      _cont=false;
                    }else{
                      _cont=true;
                    }
                  });
                },
                ),
              ),
              controller: pasw,
            ),
            RaisedButton(
              color: Colors.blue,
              child: Text('Iniciar sesion',
                  style: TextStyle(fontSize: 20,color: Colors.white)),
              onPressed: signIn,
            )
          ],
        ),
      ),
    );
  }

  void signIn() async {

    _pushScreen(context,QrScan());


  }


  void _pushScreen(BuildContext context, Widget screen) {
    Navigator.push(
      context,
      MaterialPageRoute(
          builder: (context) => screen),
    );
  }

}


