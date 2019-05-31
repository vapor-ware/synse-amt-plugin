# -*- coding: utf-8 -*-

"""Python commands for AMT power status and management."""

import sys

from amt.client import Client


def main():
    """Get or set the power state for an AMT enabled device."""

    ip = sys.argv[1]
    password = sys.argv[2]
    command_name = sys.argv[3]

    amt_client = Client(ip, password)

    commands = {
        'status': amt_client.power_status,
        'on': amt_client.power_on,
        'off': amt_client.power_off,
        'cycle': amt_client.power_cycle
    }

    # See: https://www.dmtf.org/sites/default/files/standards/documents/DSP1027_2.0.0.pdf
    # Section 7.3.1, Table 3
    status_codes = {
        '2': 'on',
        '8': 'off'
    }

    if command_name in commands:
        output = commands.get(command_name)()
        if command_name == 'status':
            output = status_codes.get(output)
            if not output:
                raise ValueError('Unexpected status "{}" returned for amt power status', output)
            sys.stdout.write(output)
    else:
        raise ValueError('No command {} found for amt power device'.format(command_name))


if __name__ == '__main__':
    main()
