# -*- coding: utf-8 -*-

"""Python commands for AMT boot target selection."""

import sys

from amt.client import Client


def main():
    """Set the boot target for an AMT-enabled device."""

    host = sys.argv[1]
    password = sys.argv[2]
    boot_target = sys.argv[3]

    amt_client = Client(host, password)
    amt_client.set_next_boot(boot_target)
    

if __name__ == '__main__':
    main()
