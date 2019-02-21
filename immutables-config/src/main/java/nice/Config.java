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
