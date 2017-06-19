import time
from indicatif.progress import ProgressBar, ProgressStyle


def main():
    downloaded = 0
    total_size = 231231231

    ps = ProgressStyle()
    ps.set_progress_chars("#>-")
    ps.set_template("{spinner:.green} [{elapsed_precise}] "
                    "[{bar:40.cyan/blue}] {bytes}/{total_bytes} "
                    "({eta})")

    pb = ProgressBar(total_size)
    pb.set_style(ps)

    while downloaded < total_size:
        new = min(downloaded + 223211, total_size)
        downloaded = new

        pb.set_position(new)
        time.sleep(12 / 1000)

    pb.finish_with_message("downloaded")


if __name__ == '__main__':
    main()
