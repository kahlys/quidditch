class Player {
  final int id;
  final String firstName;
  final String lastName;
  final String country;
  final String role;
  final int stamina;
  final int power;

  Player({
    required this.id,
    required this.firstName,
    required this.lastName,
    required this.country,
    required this.role,
    required this.stamina,
    required this.power,
  });

  factory Player.fromJson(Map<String, dynamic> json) {
    return Player(
      id: json['id'],
      firstName: json['first_name'],
      lastName: json['last_name'],
      country: json['country'],
      role: json['role'],
      stamina: json['stamina'],
      power: json['power'],
    );
  }
}
