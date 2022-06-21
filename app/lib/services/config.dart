import 'package:flutter/foundation.dart' show kIsWeb, kDebugMode;

class ConfigService {
  static String get backendHost {
    if (kDebugMode) {
      return 'localhost';
    }
    if (kIsWeb) {
      return Uri.base.host;
    }
    throw Exception("not implemented for a platform other than web");
  }

  static int get grpcWebPort {
    if (kDebugMode) {
      return 8080;
    }
    if (kIsWeb) {
      return Uri.base.port;
    }
    throw Exception("not implemented for a platform other than web");
  }

  static bool get httpsEnabled {
    if (kIsWeb) {
      if (!Uri.base.hasScheme) {
        return true;
      }
      return Uri.base.isScheme("HTTPS");
    }
    throw Exception("not implemented for a platform other than web");
  }

  static bool get todayIsTheDay {
    DateTime now = DateTime.now();
    DateTime fromDate = DateTime(2022, 6, 24, 7, 0, 0);
    DateTime toDate = DateTime(2022, 6, 29, 7, 0, 0);

    return (now.year >= fromDate.year &&
        now.year <= toDate.year &&
        now.month >= fromDate.month &&
        now.month <= toDate.month &&
        now.day >= fromDate.day &&
        now.day <= toDate.day);
  }
}
