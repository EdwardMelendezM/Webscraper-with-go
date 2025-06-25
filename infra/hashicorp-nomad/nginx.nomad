job "web" {
  datacenters = ["dc1"]

  group "nginx-group" {
    count = 2

    task "nginx" {
      driver = "docker"

      config {
        image = "nginx:latest"
        port_map {
          http = 80
        }
      }

      resources {
        network {
          port "http" {
            static = 8080
          }
        }
      }
    }
  }
}

# Run the job
# nomad run nginx.nomad