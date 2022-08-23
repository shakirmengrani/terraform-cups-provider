terraform {
  required_providers {
    cloud = {
      source  = "github.com/shakirmengrani/cloud"
    }
  }
}

provider "cloud" {
   username = "shakirmengrani"
   token = "shakirmengrani"
}

resource "cloud_order" "mycart" {
  items {
    name = "Item1"
    rate = 100
    qty = 1
  }
}