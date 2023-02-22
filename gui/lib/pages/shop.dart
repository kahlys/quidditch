import 'package:flutter/material.dart';

import 'package:gui/widgets/drawer.dart';

class PageShop extends StatefulWidget {
  @override
  _PageShop createState() => _PageShop();
}

class _PageShop extends State<PageShop> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Upgrade your team'),
      ),
      drawer: const AppDrawer(),
      body: const Center(
        child: Text(
          'Welcome to the moon',
          style: TextStyle(fontSize: 24),
        ),
      ),
    );
  }
}
