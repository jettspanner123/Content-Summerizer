function getGreeting(name) {
  const salutation = "Hello";
  const separator = ", ";
  const completeMessage = salutation + separator + name + "!";
  return completeMessage;
}
const targetPerson = "World";
const greetingText = getGreeting(targetPerson);
console.log(greetingText);