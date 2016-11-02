YAML = require('yamljs');

// parse YAML string
nativeObject = YAML.load("file.yaml");

console.log(nativeObject);
console.log(nativeObject.mood.name);
