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
