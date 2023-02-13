import 'package:flutter/material.dart';
import 'package:gui/api.dart';
import 'package:gui/models/player.dart';
import 'package:gui/widgets/drawer.dart';

class PageShop extends StatefulWidget {
  const PageShop({super.key});

  @override
  _PageShop createState() => _PageShop();
}

class _PageShop extends State<PageShop> {
  final DioClient _client = DioClient();
  List<Player> _players = [];

  int _selectedPlayerId = -1;

  @override
  Widget build(BuildContext context) {
     return Scaffold(
        appBar: AppBar(
          title: const Text('Upgrade your team'),
        ),
        drawer: const AppDrawer(),
        body: SafeArea(
          child: SingleChildScrollView(
            child: _buildPlayersTable(),
          ),
        ),
    );
  }

  @override
  void initState() {
    super.initState();
    _client.fetchPlayers().then((players) {
      setState(() {
        _players = players;
      });
    }).catchError((error) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
          content: Text('Error: $error'),
          duration: const Duration(seconds: 5),
          backgroundColor: Colors.red,
        ),
      );
    });
  }

  Widget _buildPlayersTable() {
    return Center(
      child: _players.isNotEmpty
          ? DataTable(
              columns: const [
                DataColumn(label: Text('Name')),
                DataColumn(label: Text('Country')),
                DataColumn(label: Text('Role')),
                DataColumn(label: Text('Stamina')),
                DataColumn(label: Text('Power')),
                DataColumn(label: Text(' ')),
              ],
              rows: _players
                  .map(
                    (player) => DataRow(
                      cells: [
                        DataCell(
                            Text('${player.firstName} ${player.lastName}')),
                        DataCell(Text(player.country)),
                        DataCell(Text(player.role)),
                        DataCell(Text('${player.stamina}')),
                        DataCell(Text('${player.power}')),
                        DataCell(
                          ElevatedButton(
                            child: const Text("recruit"),
                            onPressed: () {
                              print("hhh");
                              setState(() {
                                _selectedPlayerId = player.id;
                              });
                            },
                          ),
                        ),
                      ],
                    ),
                  )
                  .toList(),
            )
          : const CircularProgressIndicator(),
    );
  }
}
