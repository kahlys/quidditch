import 'package:flutter/material.dart';

class PageShop extends StatelessWidget {
  const PageShop({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Page 1'),
      ),
      drawer: Drawer(
         child: ListView(
            padding: EdgeInsets.zero,
            children: <Widget>[
              const DrawerHeader(
                decoration: BoxDecoration(
                  color: Colors.blue,
                ),
                child: Text('Menu'),
              ),
              ListTile(
                title: const Text('Shop'),
                onTap: () {
                  Navigator.pushNamed(context, '/shop');
                },
              ),
            ],
          ),
       ),
      body: const Center(
        child: Text(
          'Welcome to the moon',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}