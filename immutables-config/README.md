# Immutables & Config

There are two great libraries:

* [Immutables](https://imutables.github.io) - compensates lack of meaningful way
to have data objects in Java without adhering to rudimentary JavaBean way of having
getters and setters
* [Typesafe config](https://github.com/lightbend/config) - helps to avoid manual config
file reading (which in Java is surprisingly difficult) and type casting

## Config library restriction

Unfortunately config requires object that is auto-populated to comply to rudimentary
JavaBean spec:

```java
// ugly/Config.java

package ugly;

import java.util.List;

public class Config {
  private List<String> libraries;

  public List<String> getLibraries() {
    return libraries;
  }

  public void setLibraries(final List<String> libraries) {
    this.libraries = libraries;
  }
}
```

Then config could be read by:

```java
final Config config = ConfigFactory.load();
final ugly.Config uglyConfig = ConfigBeanFactory.create(config, ugly.Config.class);
```

See [LoadUglyConfig.java](src/test/java/LoadUglyConfig.java)

Question is: can we have the same nice auto-load thing but with nice immutable config object?

## Immutables example

Immutables allow us to define data class as interface with `@Immutables` annotation:

```java
package nice;

import java.util.List;

import org.immutables.value.Value;

@Value.Immutable
public interface Config {
  List<String> libraries();
}
```

Then we should compile code (e.g. run `compileJava` gradle task) and immutables
would generate implementation of the interface in `build/generated/sources/annotationProcessor/java/main/nice/ImmutableConfig.java`.
It could be used in a following way:

```java
final Config config = ImmutableConfig.builder()
    .addLibraries("immutables", "config")
    .build();
```

See [LoadNiceConfig.java](src/test/java/LoadNiceConfig.java)