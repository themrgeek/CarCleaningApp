import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  Future<Map<String, dynamic>> fetchUserProfile() async {
    final response =
        await http.get(Uri.parse('http://localhost:8080/user/profile'));
    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to load profile');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Profile Page'),
      ),
      body: Center(
        child: FutureBuilder<Map<String, dynamic>>(
          future: fetchUserProfile(),
          builder: (context, snapshot) {
            if (snapshot.connectionState == ConnectionState.waiting) {
              return const CircularProgressIndicator();
            } else {
              final profile = snapshot.data!;
              final user = profile['user'];
              final bookings = profile['bookings'] as List;

              return Column(
                children: [
                  const Icon(Icons.account_circle, size: 100),
                  const SizedBox(height: 20),
                  Text(
                    'User Name: ${user['name']}',
                    style: const TextStyle(fontSize: 20),
                  ),
                  Text(
                    'Email: ${user['email']}',
                    style: const TextStyle(fontSize: 20),
                  ),
                  const SizedBox(height: 20),
                  const Text(
                    'Bookings:',
                    style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                  ),
                  Expanded(
                    child: ListView.builder(
                      itemCount: bookings.length,
                      itemBuilder: (context, index) {
                        final booking = bookings[index];
                        return ListTile(
                          title: Text(booking['details']),
                          subtitle: Text('Review: ${booking['review']}'),
                        );
                      },
                    ),
                  ),
                ],
              );
            }
          },
        ),
      ),
    );
  }
}
