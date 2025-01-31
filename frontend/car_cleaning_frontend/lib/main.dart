import 'package:flutter/material.dart';
import 'splash_screen.dart'; // Import the splash screen
import 'auth/login_page.dart';
import 'package:car_cleaning_frontend/content/settings_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'AARON CLEANING SERVICES',
      theme: ThemeData(primarySwatch: Colors.yellow),
      home: const SplashScreen(), // Set SplashScreen as the home
      routes: {
        '/login': (context) => const LoginPage(),
        '/settings': (context) => const SettingsPage(),
      },
    );
  }
}
