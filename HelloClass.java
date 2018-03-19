

interface LambdaFunction {
  void call();
}

public class HelloClass {

  public static void main(String[] args) {

    // Lambda Expression
    LambdaFunction helloLambda = () -> { System.out.println("Hello World"); };

    helloLambda.call();
  }
}
