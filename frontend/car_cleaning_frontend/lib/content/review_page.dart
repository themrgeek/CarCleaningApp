// review_page.dart
import 'package:flutter/material.dart';

class ReviewPage extends StatefulWidget {
  const ReviewPage({super.key});

  @override
  State<ReviewPage> createState() => _ReviewPageState();
}

class _ReviewPageState extends State<ReviewPage> {
  final TextEditingController _reviewController = TextEditingController();
  double _rating = 3.0;

  void _submitReview() {
    // Handle review submission logic here
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(content: Text('Review Submitted!')),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Review Cleaner'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            const Text(
              'Rate the cleaner:',
              style: TextStyle(fontSize: 18),
            ),
            Slider(
              value: _rating,
              min: 1,
              max: 5,
              divisions: 4,
              label: _rating.round().toString(),
              onChanged: (double value) {
                setState(() {
                  _rating = value;
                });
              },
            ),
            const SizedBox(height: 10),
            TextField(
              controller: _reviewController,
              decoration: const InputDecoration(
                labelText: 'Write a review',
                border: OutlineInputBorder(),
              ),
              maxLines: 3,
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: _submitReview,
              child: const Text('Submit Review'),
            ),
          ],
        ),
      ),
    );
  }
}
