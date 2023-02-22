import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:gui/models/player.dart';

class DioClient {
  final Dio _dio = Dio();

  final _baseUrl = 'http://localhost:8080/api';

  Future? login(
      {required String name,
      required String email,
      required String password}) async {
    try {
      Response userData = await _dio.post(
        '$_baseUrl/login',
        data: jsonEncode({
          "name": name,
          "email": email,
          "password": password,
        }),
      );
      print('User Info: ${userData.data}');
    } on DioError catch (e) {
      if (e.response != null) {
        print('Dio error!');
        print('STATUS: ${e.response?.statusCode}');
        print('DATA: ${e.response?.data}');
        print('HEADERS: ${e.response?.headers}');
      } else {
        print('Error sending request!');
        print(e.message);
      }
    }
  }

  Future<List<Player>> fetchPlayers() async {
    final response = await _dio.get('$_baseUrl/shop/players');
    if (response.statusCode == 200) {
      final data = json.decode(response.data) as Map<String, dynamic>;
      final players = data['players'] as List<dynamic>;
      return players.map((player) => Player.fromJson(player)).toList();
    } else {
      throw Exception('Failed to fetch players');
    }
  }
}
