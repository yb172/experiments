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
// nice/Config.java
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

## Problem

Unfortunately if we would try to combine two objects:
`ConfigBeanFactory.create(config, nice.Config.class)` then we would get an exception:

```text
com.typesafe.config.ConfigException$BadBean: nice.Config needs a public no-args constructor to be used as a bean
```

Ok, it makes sense: we're passing interface which doesn't have constructor.
But `ImmutableConfig` doesn't have constructor either - it has builder class.

So maybe we could pass builder class?

```java
final ImmutableConfig.Builder builder = ConfigBeanFactory.create(config, ImmutableConfig.Builder.class);
final nice.Config niceConfig = builder.build();
```

Well, in that case we would get another problem:

```text
com.typesafe.config.ConfigException$BadBean: nice.ImmutableConfig$Builder getters and setters are not accessible, they must be for use as a bean
```

That makes sense, our builder by default generates method names which are the same as interface method names.

## Solution

To solve the problem we could use `@Modifiable` annotation instead of `@Immutable`
as [suggested in issue](https://immutables.github.io/immutable.html#plain-public-constructor):

```java
// nice/Config.java
package nice;

import java.util.List;

import org.immutables.value.Value;

@Value.Modifiable
@Value.Style(
    create = "new", // rename create method to "new", turning factory into plain public constuctor,
    beanFriendlyModifiables = true // setters will return void instead of instance if needed for strict javabean framework
)
public interface Config {
  List<String> getLibraries();
}
```

Then our config loading code would look like:

```java
final com.typesafe.config.Config config = ConfigFactory.load();
final nice.Config niceConfig = ConfigBeanFactory.create(config, ModifiableConfig.class);
```

And it would work perfectly fine.

See [LoadNiceConfig.java](src/test/java/LoadNiceConfig.java)
