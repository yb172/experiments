import org.assertj.core.api.WithAssertions;
import org.junit.Test;
import nice.Config;
import nice.ImmutableConfig;

public class LoadNiceConfig implements WithAssertions {

  @Test
  public void exampleOfImmutableObject() {
    final Config config = ImmutableConfig.builder()
        .addLibraries("immutables", "config")
        .build();
    assertThat(config.libraries()).containsExactly("immutables", "config");
  }
}
