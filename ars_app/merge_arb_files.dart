import 'dart:convert';
import 'dart:io';

// Output file path. Do not change this.
const List<String> directoryPaths = ['./lib/screen', './lib/base'];
const String enOutputFilePath = './lib/l10n/app_en.arb';
const String viOutputFilePath = './lib/l10n/app_vi.arb';

Future<Map<String, dynamic>> readFile(String filePath) async {
  var input = await File(filePath).readAsString();
  var map = jsonDecode(input);
  return map;
}

Future<void> writeFile(String filePath, Map<String, dynamic> data) async {
  await File(filePath).writeAsString(jsonEncode(data));
}

Future<Map<String, dynamic>> mergeMap(List<String> filePaths) async {
  var map = <String, dynamic>{};
  for (final filePath in filePaths) {
    var newMap = await readFile(filePath);
    map.addAll(newMap);
  }
  return map;
}

Future<List<File>> getArbFiles(String directoryPath) async {
  final directory = Directory(directoryPath);
  if (!directory.existsSync()) {
    throw FileSystemException('Directory does not exist: $directoryPath');
  }

  final arbFiles = <File>[];

  await for (final entity in directory.list()) {
    if (entity is File && entity.path.endsWith('.arb')) {
      arbFiles.add(entity);
    } else if (entity is Directory) {
      final subDirectoryPath = entity.path;
      final subDirectoryArbFiles = await getArbFiles(subDirectoryPath);
      arbFiles.addAll(subDirectoryArbFiles);
    }
  }

  return arbFiles;
}

void main() async {
  List<String> viFilePath = [];
  List<String> enFilePath = [];

  for (final directoryPath in directoryPaths) {
    final arbFiles = await getArbFiles(directoryPath);
    for (final file in arbFiles) {
      if (file.path.endsWith('_vi.arb')) {
        viFilePath.add(file.path);
      } else if (file.path.endsWith('_en.arb')) {
        enFilePath.add(file.path);
      }
    }
  }

  var enMap = await mergeMap(enFilePath);
  await writeFile(enOutputFilePath, enMap);

  var viMap = await mergeMap(viFilePath);
  await writeFile(viOutputFilePath, viMap);
}
