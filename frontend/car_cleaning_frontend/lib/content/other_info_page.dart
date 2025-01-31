import 'package:flutter/material.dart';

class OtherInfoPage extends StatelessWidget {
  const OtherInfoPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Other Info'),
      ),
      body: const Center(
        child: Text('This is the Other Info page'),
      ),
    );
  }
}
