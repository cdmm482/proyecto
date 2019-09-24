import 'package:flutter/material.dart';
import 'package:hola/page/qr_scan.dart';

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
      home: QrScan(),
    );
  }
}