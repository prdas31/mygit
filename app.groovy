@RestController
class Helloworld {

    @RequestMapping("/")
    String home() {
        "Hello, world!"
    }
}
