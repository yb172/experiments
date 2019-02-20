import com.typesafe.config.ConfigBeanFactory;
import com.typesafe.config.ConfigFactory;
import org.assertj.core.api.WithAssertions;
import org.junit.Test;
import nice.ModifiableConfig;

public class LoadNiceConfig implements WithAssertions {

  /* Tests for nice.Config annotated with @Value.Immutable

  @Test
  public void exampleOfImmutableObject() {
    final Config config = ImmutableConfig.builder()
        .addLibraries("immutables", "config")
        .build();
    assertThat(config.getLibraries()).containsExactly("immutables", "config");
  }


  @Test
  public void tryLoadInterface() {
    final com.typesafe.config.Config config = ConfigFactory.load();
    final nice.Config niceConfig = ConfigBeanFactory.create(config, nice.Config.class);
    assertThat(niceConfig.getLibraries()).containsExactly("immutables", "config");
    // com.typesafe.config.ConfigException$BadBean: nice.Config needs a public no-args constructor to be used as a bean
    // exception whould be thrown
  }
  */

  @Test
  public void tryLoadBuilder() {
    final com.typesafe.config.Config config = ConfigFactory.load();
    final nice.Config niceConfig = ConfigBeanFactory.create(config, ModifiableConfig.class);
    assertThat(niceConfig.getLibraries()).containsExactly("immutables", "config");
  }
}
