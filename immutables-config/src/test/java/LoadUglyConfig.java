import com.typesafe.config.Config;
import com.typesafe.config.ConfigBeanFactory;
import com.typesafe.config.ConfigFactory;
import org.assertj.core.api.WithAssertions;
import org.junit.Test;

public class LoadUglyConfig implements WithAssertions {

  @Test
  public void test() {
    final Config config = ConfigFactory.load();
    final ugly.Config uglyConfig = ConfigBeanFactory.create(config, ugly.Config.class);
    assertThat(uglyConfig.getLibraries()).containsExactly("immutables", "config");
  }
}
