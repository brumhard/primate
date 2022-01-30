import 'package:flutter/foundation.dart' show kIsWeb;

class ConfigService {
  static String get backendHost {
    if (kIsWeb) {
      return Uri.base.host;
    }
    throw Exception("not implemented for a platform other than web");
  }

  static int get grpcWebPort {
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
}
