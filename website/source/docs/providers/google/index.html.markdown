---
layout: "google"
page_title: "Provider: Google Cloud"
sidebar_current: "docs-google-index"
---

# Google Cloud Provider

The Google Cloud provider is used to interact with
[Google Cloud services](https://cloud.google.com/). The provider needs
to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
# Configure the Google Cloud provider
provider "google" {
    account_file = "account.json"
    project = "my-gce-project"
    region = "us-central1"
}

# Create a new instance
resource "google_compute_instance" "default" {
    ...
}
```

## Configuration Reference

The following keys can be used to configure the provider.

* `account_file` - (Required) Path to the JSON file used to describe
  your account credentials, downloaded from Google Cloud Console. More
  details on retrieving this file are below.

* `project` - (Required) The name of the project to apply any resources to.

* `region` - (Required) The region to operate under.

## Authentication JSON File

Authenticating with Google Cloud services requires a JSON
file which we call the _account file_.

This file is downloaded directly from the
[Google Developers Console](https://console.developers.google.com). To make
the process more straightforwarded, it is documented here:

1. Log into the [Google Developers Console](https://console.developers.google.com)
   and select a project.

2. Under the "APIs & Auth" section, click "Credentials."

3. Create a new OAuth client ID and select "Service account" as the type
   of account. Once created, a JSON file should be downloaded. This is your
   _account file_.
