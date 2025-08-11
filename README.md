# AWS EC2 Instance Count

This Go script is designed to efficiently count the number of running EC2 instances across multiple AWS accounts and regions. It serves as a practical demonstration of how **Go's native concurrency tools**—specifically `goroutines` and `sync.WaitGroup`—can provide dramatic performance improvements over a traditional synchronous approach.

### Features

* **Multi-Account Support:** The script automatically detects and iterates through all AWS profiles configured in your `~/.aws/config` file, which is perfect for environments using AWS IAM Identity Center (SSO).
* **Multi-Region Support:** It counts instances across a predefined list of AWS regions.
* **Performance Comparison:** It runs the same task both synchronously and concurrently, providing a direct, side-by-side comparison of execution times.
* **Lightweight & Fast:** The concurrent version is highly performant, capable of executing dozens of API calls in parallel to deliver results in a fraction of the time.

### Prerequisites

To run this script, you'll need the following installed and configured on your system:

1.  **Go:** Version 1.16 or newer.
2.  **AWS CLI:** The AWS Command Line Interface should be installed.
3.  **AWS Credentials:** Your `~/.aws/config` file must be set up with the necessary profiles and credentials, typically configured for AWS SSO.

### Setup and Installation

1.  Save the code to a file named `main.go`.
2.  Open your terminal in the same directory and initialize a Go module.
    ```bash
    go mod init ec2-instance-counter
    ```
3.  Install the required AWS SDK and `ini` parsing libraries.
    ```bash
    go get [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2/config)
    go get [github.com/aws/aws-sdk-go-v2/service/ec2](https://github.com/aws/aws-sdk-go-v2/service/ec2)
    go get gopkg.in/ini.v1
    ```

---
### How It Works

#### Synchronous Approach

This is the traditional, sequential method. The script uses nested `for` loops to process each AWS account and region one by one. It must complete an API call for the current region and profile before moving on to the next. The total execution time is the sum of all individual API call latencies.

```go
for _, profile := range profiles {
    for _, region := range regions {
        if err := listEC2Instances(region, profile); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}
```
# AWS EC2 Instance Count

This Go script is designed to efficiently count the number of running EC2 instances across multiple AWS accounts and regions. It serves as a practical demonstration of how **Go's native concurrency tools**—specifically `goroutines` and `sync.WaitGroup`—can provide dramatic performance improvements over a traditional synchronous approach.

### Features

* **Multi-Account Support:** The script automatically detects and iterates through all AWS profiles configured in your `~/.aws/config` file, which is perfect for environments using AWS IAM Identity Center (SSO).
* **Multi-Region Support:** It counts instances across a predefined list of AWS regions.
* **Performance Comparison:** It runs the same task both synchronously and concurrently, providing a direct, side-by-side comparison of execution times.
* **Lightweight & Fast:** The concurrent version is highly performant, capable of executing dozens of API calls in parallel to deliver results in a fraction of the time.

### Prerequisites

To run this script, you'll need the following installed and configured on your system:

1.  **Go:** Version 1.16 or newer.
2.  **AWS CLI:** The AWS Command Line Interface should be installed.
3.  **AWS Credentials:** Your `~/.aws/config` file must be set up with the necessary profiles and credentials, typically configured for AWS SSO.

### Setup and Installation

1.  Save the code to a file named `main.go`.
2.  Open your terminal in the same directory and initialize a Go module.
    ```bash
    go mod init ec2-instance-counter
    ```
3.  Install the required AWS SDK and `ini` parsing libraries.
    ```bash
    go get [github.com/aws/aws-sdk-go-v2/config](https://github.com/aws/aws-sdk-go-v2/config)
    go get [github.com/aws/aws-sdk-go-v2/service/ec2](https://github.com/aws/aws-sdk-go-v2/service/ec2)
    go get gopkg.in/ini.v1
    ```

---
### How It Works

#### Synchronous Approach

This is the traditional, sequential method. The script uses nested `for` loops to process each AWS account and region one by one. It must complete an API call for the current region and profile before moving on to the next. The total execution time is the sum of all individual API call latencies.

```go
for _, profile := range profiles {
    for _, region := range regions {
        if err := listEC2Instances(region, profile); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
}
```