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

  @Override
  public String toString() {
    return "Config{" +
        "getLibraries=" + libraries +
        '}';
  }
}
